package types

type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 User    `json:"creator"`
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name,omitempty"`
	ExpireDate              *int64  `json:"expire_date,omitempty"`
	MemberLimit             *int64  `json:"member_limit,omitempty"`
	PendingJoinRequestCount *int64  `json:"pending_join_request_count,omitempty"`
}
