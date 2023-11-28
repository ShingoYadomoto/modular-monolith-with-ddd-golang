package domain

import (
	"time"

	"github.com/google/uuid"
)

/*
using System;
using MediatR;

namespace CompanyName.MyMeetings.BuildingBlocks.Domain
{
    public interface IDomainEvent : INotification
    {
        Guid Id { get; }

        DateTime OccurredOn { get; }
    }
}
*/

type IDomainEvent interface {
	ID() uuid.UUID
	OccurredOn() time.Time
}
