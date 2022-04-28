package types

type ChatMembers []ChatMember

func (cm ChatMembers) Owners() []ChatMemberOwner {
	var res []ChatMemberOwner
	for _, member := range cm {
		if o := member.Owner(); o != nil {
			res = append(res, *o)
		}
	}
	return res
}

func (cm ChatMembers) Administrators() []ChatMemberAdministrator {
	var res []ChatMemberAdministrator
	for _, member := range cm {
		if o := member.Administrator(); o != nil {
			res = append(res, *o)
		}
	}
	return res
}

func (cm ChatMembers) Members() []ChatMemberMember {
	var res []ChatMemberMember
	for _, member := range cm {
		if o := member.Member(); o != nil {
			res = append(res, *o)
		}
	}
	return res
}

func (cm ChatMembers) Restricteds() []ChatMemberRestricted {
	var res []ChatMemberRestricted
	for _, member := range cm {
		if o := member.Restricted(); o != nil {
			res = append(res, *o)
		}
	}
	return res
}
func (cm ChatMembers) Lefts() []ChatMemberLeft {
	var res []ChatMemberLeft
	for _, member := range cm {
		if o := member.Left(); o != nil {
			res = append(res, *o)
		}
	}
	return res
}

func (cm ChatMembers) Banneds() []ChatMemberBanned {
	var res []ChatMemberBanned
	for _, member := range cm {
		if o := member.Banned(); o != nil {
			res = append(res, *o)
		}
	}
	return res
}
