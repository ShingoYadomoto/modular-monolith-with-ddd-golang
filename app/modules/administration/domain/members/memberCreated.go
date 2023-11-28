package members

import (
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
)

/*
using CompanyName.MyMeetings.BuildingBlocks.Domain;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.Members.Events
{
    public class MemberCreatedDomainEvent : DomainEventBase
    {
        public MemberCreatedDomainEvent(MemberId memberId)
        {
            MemberId = memberId;
        }

        public MemberId MemberId { get; }
    }
}
*/

type MemberCreatedDomainEvent struct {
	MemberID MemberID
	domain.DomainEventBase
}

func NewMemberCreatedDomainEvent(memberID MemberID) *MemberCreatedDomainEvent {
	return &MemberCreatedDomainEvent{
		MemberID:        memberID,
		DomainEventBase: *domain.NewDomainEventBase(),
	}
}
