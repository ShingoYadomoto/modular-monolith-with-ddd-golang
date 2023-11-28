package members

import (
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
	"github.com/google/uuid"
)

/*
using System;
using CompanyName.MyMeetings.BuildingBlocks.Domain;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.Members
{
    public class MemberId : TypedIdValueBase
    {
        public MemberId(Guid value)
            : base(value)
        {
        }
    }
}
*/

type MemberID struct {
	domain.TypedIDValueBase
}

func NewMemberID(value uuid.UUID) MemberID {
	return MemberID{domain.NewTypedIDValueBase(value)}
}
