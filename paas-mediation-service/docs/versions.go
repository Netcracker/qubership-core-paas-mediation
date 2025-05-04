package docs

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/netcracker/qubership-core-lib-go/v3/logging"
	"regexp"
	"sort"
	"strconv"
)

var (
	MajorVersion         int
	MinorVersion         int
	SupportedMajors      []int
	logger               logging.Logger
	pathWithVersionRegex = regexp.MustCompile(`^.*(/v(\d+)).*$`)
	sinceTagRegex        = regexp.MustCompile(`^since:(\d+)\.(\d+)$`)
)

func init() {
	logger = logging.GetLogger("versions")
	var err error
	MajorVersion, MinorVersion, SupportedMajors, err = getVersions()
	if err != nil {
		panic(err.Error())
	}
	SwaggerInfo.Version = fmt.Sprintf("%d.%d", MajorVersion, MinorVersion)
	logger.Info("Versions -> Major:%d Minor:%d SupportedMajors:%v", MajorVersion, MinorVersion, SupportedMajors)
}

func getVersions() (major, minor int, majors []int, err error) {
	var swaggerSpec spec.Swagger
	err = json.Unmarshal([]byte(SwaggerInfo.ReadDoc()), &swaggerSpec)
	if err != nil {
		err = fmt.Errorf("failed to resolve versions: %w", err)
		return
	}
	versionsMap := map[int]int{}
	for path, item := range swaggerSpec.Paths.Paths {
		// check path's version and majorVersion from 'since' tag
		var versionFromPathInt int
		if pathWithVersionRegex.MatchString(path) {
			versionFromPath := pathWithVersionRegex.FindStringSubmatch(path)[2]
			versionFromPathInt, _ = strconv.Atoi(versionFromPath)
		} else {
			// path without api version (like /health, /api-version etc), skip
			continue
		}
		for _, operation := range []*spec.Operation{item.Get, item.Put, item.Post, item.Delete, item.Options, item.Head, item.Patch} {
			if operation == nil {
				continue
			}
			var pathMajorVersion int
			var pathMinorVersion int
			pathMajorVersion, pathMinorVersion, err = parseSinceTag(operation.Tags)
			if err != nil {
				return
			}
			if versionFromPathInt != pathMajorVersion {
				err = fmt.Errorf("version '%d' from path '%s', and major version '%d' from 'since' tag '%s' must be the same",
					versionFromPathInt, path, pathMajorVersion, getSinceTag(operation.Tags))
				return
			}
			if minorVersion, ok := versionsMap[pathMajorVersion]; ok {
				if minorVersion < pathMinorVersion {
					versionsMap[pathMajorVersion] = pathMinorVersion
				}
			} else {
				versionsMap[pathMajorVersion] = pathMinorVersion
			}
		}
	}
	for majorVersion := range versionsMap {
		majors = append(majors, majorVersion)
	}
	sort.Ints(majors)
	major = majors[len(majors)-1]
	minor = versionsMap[major]
	return
}

func getSinceTag(tags []string) string {
	for _, tag := range tags {
		if sinceTagRegex.MatchString(tag) {
			return tag
		}
	}
	return ""
}

func parseSinceTag(tags []string) (major, minor int, err error) {
	tag := getSinceTag(tags)
	if tag == "" {
		err = fmt.Errorf("invalid or missing 'since' tag. Must match regex = %s", sinceTagRegex.String())
	} else {
		sinceTagMatch := sinceTagRegex.FindStringSubmatch(tag)
		major, _ = strconv.Atoi(sinceTagMatch[1])
		minor, _ = strconv.Atoi(sinceTagMatch[2])
	}
	return
}
