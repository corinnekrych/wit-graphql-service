// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-wit/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit
// --version=v1.3.0

package app

import (
	"fmt"
	"strings"
)

// UserServiceHref returns the resource href.
func UserServiceHref() string {
	return "/api/user/services"
}

// AreaHref returns the resource href.
func AreaHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/areas/%v", paramid)
}

// CodebaseHref returns the resource href.
func CodebaseHref(codebaseID interface{}) string {
	paramcodebaseID := strings.TrimLeftFunc(fmt.Sprintf("%v", codebaseID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/codebases/%v", paramcodebaseID)
}

// CommentsHref returns the resource href.
func CommentsHref(commentID interface{}) string {
	paramcommentID := strings.TrimLeftFunc(fmt.Sprintf("%v", commentID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/comments/%v", paramcommentID)
}

// FeaturesHref returns the resource href.
func FeaturesHref(featureName interface{}) string {
	paramfeatureName := strings.TrimLeftFunc(fmt.Sprintf("%v", featureName), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/features/%v", paramfeatureName)
}

// FilterHref returns the resource href.
func FilterHref() string {
	return "/api/filters"
}

// IterationHref returns the resource href.
func IterationHref(iterationID interface{}) string {
	paramiterationID := strings.TrimLeftFunc(fmt.Sprintf("%v", iterationID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/iterations/%v", paramiterationID)
}

// LabelHref returns the resource href.
func LabelHref(spaceID, labelID interface{}) string {
	paramspaceID := strings.TrimLeftFunc(fmt.Sprintf("%v", spaceID), func(r rune) bool { return r == '/' })
	paramlabelID := strings.TrimLeftFunc(fmt.Sprintf("%v", labelID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/spaces/%v/labels/%v", paramspaceID, paramlabelID)
}

// NamedWorkItemsHref returns the resource href.
func NamedWorkItemsHref(userName, spaceName, wiNumber interface{}) string {
	paramuserName := strings.TrimLeftFunc(fmt.Sprintf("%v", userName), func(r rune) bool { return r == '/' })
	paramspaceName := strings.TrimLeftFunc(fmt.Sprintf("%v", spaceName), func(r rune) bool { return r == '/' })
	paramwiNumber := strings.TrimLeftFunc(fmt.Sprintf("%v", wiNumber), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/namedspaces/%v/%v/workitems/%v", paramuserName, paramspaceName, paramwiNumber)
}

// NamedspacesHref returns the resource href.
func NamedspacesHref(userName, spaceName interface{}) string {
	paramuserName := strings.TrimLeftFunc(fmt.Sprintf("%v", userName), func(r rune) bool { return r == '/' })
	paramspaceName := strings.TrimLeftFunc(fmt.Sprintf("%v", spaceName), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/namedspaces/%v/%v", paramuserName, paramspaceName)
}

// QueryHref returns the resource href.
func QueryHref(spaceID, queryID interface{}) string {
	paramspaceID := strings.TrimLeftFunc(fmt.Sprintf("%v", spaceID), func(r rune) bool { return r == '/' })
	paramqueryID := strings.TrimLeftFunc(fmt.Sprintf("%v", queryID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/spaces/%v/queries/%v", paramspaceID, paramqueryID)
}

// SearchHref returns the resource href.
func SearchHref() string {
	return "/api/search"
}

// SpaceHref returns the resource href.
func SpaceHref(spaceID interface{}) string {
	paramspaceID := strings.TrimLeftFunc(fmt.Sprintf("%v", spaceID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/spaces/%v", paramspaceID)
}

// SpaceTemplateHref returns the resource href.
func SpaceTemplateHref(spaceTemplateID interface{}) string {
	paramspaceTemplateID := strings.TrimLeftFunc(fmt.Sprintf("%v", spaceTemplateID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/spacetemplates/%v", paramspaceTemplateID)
}

// StatusHref returns the resource href.
func StatusHref() string {
	return "/api/status"
}

// TrackerHref returns the resource href.
func TrackerHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/trackers/%v", paramid)
}

// TrackerqueryHref returns the resource href.
func TrackerqueryHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/trackerqueries/%v", paramid)
}

// UserHref returns the resource href.
func UserHref() string {
	return "/api/user"
}

// UsersHref returns the resource href.
func UsersHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/users/%v", paramid)
}

// WorkItemBoardHref returns the resource href.
func WorkItemBoardHref(boardID interface{}) string {
	paramboardID := strings.TrimLeftFunc(fmt.Sprintf("%v", boardID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/workitemboards/%v", paramboardID)
}

// WorkItemLinkHref returns the resource href.
func WorkItemLinkHref(linkID interface{}) string {
	paramlinkID := strings.TrimLeftFunc(fmt.Sprintf("%v", linkID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/workitemlinks/%v", paramlinkID)
}

// WorkItemLinkTypeHref returns the resource href.
func WorkItemLinkTypeHref(wiltID interface{}) string {
	paramwiltID := strings.TrimLeftFunc(fmt.Sprintf("%v", wiltID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/workitemlinktypes/%v", paramwiltID)
}

// WorkItemTypeGroupHref returns the resource href.
func WorkItemTypeGroupHref(groupID interface{}) string {
	paramgroupID := strings.TrimLeftFunc(fmt.Sprintf("%v", groupID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/workitemtypegroups/%v", paramgroupID)
}

// WorkitemHref returns the resource href.
func WorkitemHref(wiID interface{}) string {
	paramwiID := strings.TrimLeftFunc(fmt.Sprintf("%v", wiID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/workitems/%v", paramwiID)
}

// WorkitemtypeHref returns the resource href.
func WorkitemtypeHref(witID interface{}) string {
	paramwitID := strings.TrimLeftFunc(fmt.Sprintf("%v", witID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/workitemtypes/%v", paramwitID)
}
