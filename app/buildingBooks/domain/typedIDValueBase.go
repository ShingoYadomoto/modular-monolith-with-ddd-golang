package domain

/*
using System;

namespace CompanyName.MyMeetings.BuildingBlocks.Domain
{
    public abstract class TypedIdValueBase : IEquatable<TypedIdValueBase>
    {
        public Guid Value { get; }

        protected TypedIdValueBase(Guid value)
        {
            if (value == Guid.Empty)
            {
                throw new InvalidOperationException("Id value cannot be empty!");
            }

            Value = value;
        }

        public override bool Equals(object obj)
        {
            if (ReferenceEquals(null, obj))
            {
                return false;
            }

            return obj is TypedIdValueBase other && Equals(other);
        }

        public override int GetHashCode()
        {
            return Value.GetHashCode();
        }

        public bool Equals(TypedIdValueBase other)
        {
            return this.Value == other?.Value;
        }

        public static bool operator ==(TypedIdValueBase obj1, TypedIdValueBase obj2)
        {
            if (object.Equals(obj1, null))
            {
                if (object.Equals(obj2, null))
                {
                    return true;
                }

                return false;
            }

            return obj1.Equals(obj2);
        }

        public static bool operator !=(TypedIdValueBase x, TypedIdValueBase y)
        {
            return !(x == y);
        }
    }
}
*/

import (
	"github.com/google/uuid"
)

// TypedIDValueBase is a base struct for typed identifier values.
type TypedIDValueBase struct {
	Value uuid.UUID
}

// NewTypedIDValueBase creates a new TypedIDValueBase with the given value.
func NewTypedIDValueBase(value uuid.UUID) TypedIDValueBase {
	return TypedIDValueBase{Value: value}
}

// Equals checks if two TypedIDValueBase instances are equal.
func (id TypedIDValueBase) Equals(other TypedIDValueBase) bool {
	return id.Value == other.Value
}
