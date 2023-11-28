package meetingGroupProposals

import (
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
	"github.com/google/uuid"
)

/*
using System;
using CompanyName.MyMeetings.BuildingBlocks.Domain;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.MeetingGroupProposals
{
    public class MeetingGroupProposalId : TypedIdValueBase
    {
        public MeetingGroupProposalId(Guid value)
            : base(value)
        {
        }
    }
}
*/

type MeetingGroupProposalsID struct {
	domain.TypedIDValueBase
}

func NewMeetingGroupProposalsID(value uuid.UUID) MeetingGroupProposalsID {
	return MeetingGroupProposalsID{domain.NewTypedIDValueBase(value)}
}
