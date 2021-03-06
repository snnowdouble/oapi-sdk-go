// Code generated by lark suite oapi sdk gen
package v4

type Calendar struct {
	CalendarId   string `json:"calendar_id,omitempty"`
	Summary      string `json:"summary,omitempty"`
	Description  string `json:"description,omitempty"`
	Permissions  string `json:"permissions,omitempty"`
	Color        int    `json:"color,omitempty"`
	Type         string `json:"type,omitempty"`
	SummaryAlias string `json:"summary_alias,omitempty"`
	IsDeleted    bool   `json:"is_deleted,omitempty"`
	IsThirdParty bool   `json:"is_third_party,omitempty"`
	Role         string `json:"role,omitempty"`
}
type CalendarAcl struct {
	AclId string    `json:"acl_id,omitempty"`
	Role  string    `json:"role,omitempty"`
	Scope *AclScope `json:"scope,omitempty"`
}
type CalendarEvent struct {
	EventId          string         `json:"event_id,omitempty"`
	Summary          string         `json:"summary,omitempty"`
	Description      string         `json:"description,omitempty"`
	StartTime        *TimeInfo      `json:"start_time,omitempty"`
	EndTime          *TimeInfo      `json:"end_time,omitempty"`
	Vchat            *Vchat         `json:"vchat,omitempty"`
	Visibility       string         `json:"visibility,omitempty"`
	AttendeeAbility  string         `json:"attendee_ability,omitempty"`
	FreeBusyStatus   string         `json:"free_busy_status,omitempty"`
	Location         *EventLocation `json:"location,omitempty"`
	Color            int            `json:"color,omitempty"`
	Reminders        []*Reminder    `json:"reminders,omitempty"`
	Recurrence       string         `json:"recurrence,omitempty"`
	Status           string         `json:"status,omitempty"`
	IsException      bool           `json:"is_exception,omitempty"`
	RecurringEventId string         `json:"recurring_event_id,omitempty"`
	Schemas          []*Schema      `json:"schemas,omitempty"`
}
type CalendarEventAttendee struct {
	Type            string                `json:"type,omitempty"`
	AttendeeId      string                `json:"attendee_id,omitempty"`
	RsvpStatus      string                `json:"rsvp_status,omitempty"`
	IsOptional      bool                  `json:"is_optional,omitempty"`
	IsOrganizer     bool                  `json:"is_organizer,omitempty"`
	IsExternal      bool                  `json:"is_external,omitempty"`
	DisplayName     string                `json:"display_name,omitempty"`
	ChatMembers     []*AttendeeChatMember `json:"chat_members,omitempty"`
	UserId          string                `json:"user_id,omitempty"`
	ChatId          string                `json:"chat_id,omitempty"`
	RoomId          string                `json:"room_id,omitempty"`
	ThirdPartyEmail string                `json:"third_party_email,omitempty"`
}
type Freebusy struct {
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

type AclScope struct {
	Type   string `json:"type,omitempty"`
	UserId string `json:"user_id,omitempty"`
}
type AttendeeChatMember struct {
	RsvpStatus  string `json:"rsvp_status,omitempty"`
	IsOptional  bool   `json:"is_optional,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	IsOrganizer bool   `json:"is_organizer,omitempty"`
	IsExternal  bool   `json:"is_external,omitempty"`
}
type EventLocation struct {
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
type Reminder struct {
	Minutes int `json:"minutes,omitempty"`
}
type Schema struct {
	UiName   string `json:"ui_name,omitempty"`
	UiStatus string `json:"ui_status,omitempty"`
	AppLink  string `json:"app_link,omitempty"`
}
type TimeInfo struct {
	Date      string `json:"date,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
}
type Vchat struct {
	MeetingUrl string `json:"meeting_url,omitempty"`
}

type CalendarAclListResult struct {
	Acls      []*CalendarAcl `json:"acls,omitempty"`
	HasMore   bool           `json:"has_more,omitempty"`
	PageToken string         `json:"page_token,omitempty"`
}

type CalendarCreateResult struct {
	Calendar *Calendar `json:"calendar,omitempty"`
}

type CalendarEventAttendeeBatchDeleteReqBody struct {
	AttendeeIds []string `json:"attendee_ids,omitempty"`
}

type CalendarEventAttendeeCreateReqBody struct {
	Attendees []*CalendarEventAttendee `json:"attendees,omitempty"`
}

type CalendarEventAttendeeCreateResult struct {
	Attendees []*CalendarEventAttendee `json:"attendees,omitempty"`
}

type CalendarEventAttendeeListResult struct {
	Items     []*CalendarEventAttendee `json:"items,omitempty"`
	HasMore   bool                     `json:"has_more,omitempty"`
	PageToken string                   `json:"page_token,omitempty"`
}

type CalendarEventCreateResult struct {
	Event *CalendarEvent `json:"event,omitempty"`
}

type CalendarEventGetResult struct {
	Event *CalendarEvent `json:"event,omitempty"`
}

type CalendarEventListResult struct {
	HasMore   bool             `json:"has_more,omitempty"`
	PageToken string           `json:"page_token,omitempty"`
	SyncToken string           `json:"sync_token,omitempty"`
	Items     []*CalendarEvent `json:"items,omitempty"`
}

type CalendarEventPatchResult struct {
	Event *CalendarEvent `json:"event,omitempty"`
}

type CalendarListResult struct {
	HasMore      bool        `json:"has_more,omitempty"`
	PageToken    string      `json:"page_token,omitempty"`
	SyncToken    string      `json:"sync_token,omitempty"`
	CalendarList []*Calendar `json:"calendar_list,omitempty"`
}

type CalendarPatchResult struct {
	Calendar *Calendar `json:"calendar,omitempty"`
}

type CalendarSearchReqBody struct {
	Query string `json:"query,omitempty"`
}

type CalendarSearchResult struct {
	Items     []*Calendar `json:"items,omitempty"`
	PageToken string      `json:"page_token,omitempty"`
}

type FreebusyListReqBody struct {
	TimeMin string `json:"time_min,omitempty"`
	TimeMax string `json:"time_max,omitempty"`
	UserId  string `json:"user_id,omitempty"`
	RoomId  string `json:"room_id,omitempty"`
}

type FreebusyListResult struct {
	FreebusyList []*Freebusy `json:"freebusy_list,omitempty"`
}
