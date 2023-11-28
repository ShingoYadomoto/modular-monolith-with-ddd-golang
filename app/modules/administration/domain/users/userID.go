package users

import (
	"github.com/ShingoYadomoto/anything-go/modular-monolith-with-ddd-golang/app/buildingBooks/domain"
	"github.com/google/uuid"
)

/*
using CompanyName.MyMeetings.BuildingBlocks.Domain;

namespace CompanyName.MyMeetings.Modules.Administration.Domain.Users
{
    public class UserId : TypedIdValueBase
    {
        public UserId(Guid value)
            : base(value)
        {
        }
    }

*/

type UserID struct {
	domain.TypedIDValueBase
}

func NewUserID(value uuid.UUID) UserID {
	return UserID{domain.NewTypedIDValueBase(value)}
}
