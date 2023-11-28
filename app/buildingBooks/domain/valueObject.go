package domain

import (
	"reflect"
)

/*
using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;

namespace CompanyName.MyMeetings.BuildingBlocks.Domain
{
    public abstract class ValueObject : IEquatable<ValueObject>
    {
        private List<PropertyInfo> _properties;

        private List<FieldInfo> _fields;

        public static bool operator ==(ValueObject obj1, ValueObject obj2)
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

        public static bool operator !=(ValueObject obj1, ValueObject obj2)
        {
            return !(obj1 == obj2);
        }

        public bool Equals(ValueObject obj)
        {
            return Equals(obj as object);
        }

        public override bool Equals(object obj)
        {
            if (obj == null || GetType() != obj.GetType())
            {
                return false;
            }

            return GetProperties().All(p => PropertiesAreEqual(obj, p))
                && GetFields().All(f => FieldsAreEqual(obj, f));
        }

        public override int GetHashCode()
        {
            unchecked
            {
                int hash = 17;
                foreach (var prop in GetProperties())
                {
                    var value = prop.GetValue(this, null);
                    hash = HashValue(hash, value);
                }

                foreach (var field in GetFields())
                {
                    var value = field.GetValue(this);
                    hash = HashValue(hash, value);
                }

                return hash;
            }
        }

        protected static void CheckRule(IBusinessRule rule)
        {
            if (rule.IsBroken())
            {
                throw new BusinessRuleValidationException(rule);
            }
        }

        private bool PropertiesAreEqual(object obj, PropertyInfo p)
        {
            return object.Equals(p.GetValue(this, null), p.GetValue(obj, null));
        }

        private bool FieldsAreEqual(object obj, FieldInfo f)
        {
            return object.Equals(f.GetValue(this), f.GetValue(obj));
        }

        private IEnumerable<PropertyInfo> GetProperties()
        {
            if (this._properties == null)
            {
                this._properties = GetType()
                    .GetProperties(BindingFlags.Instance | BindingFlags.Public)
                    .Where(p => p.GetCustomAttribute(typeof(IgnoreMemberAttribute)) == null)
                    .ToList();

                // Not available in Core
                // !Attribute.IsDefined(p, typeof(IgnoreMemberAttribute))).ToList();
            }

            return this._properties;
        }

        private IEnumerable<FieldInfo> GetFields()
        {
            if (this._fields == null)
            {
                this._fields = GetType().GetFields(BindingFlags.Instance | BindingFlags.Public | BindingFlags.NonPublic)
                    .Where(p => p.GetCustomAttribute(typeof(IgnoreMemberAttribute)) == null)
                    .ToList();
            }

            return this._fields;
        }

        private int HashValue(int seed, object value)
        {
            var currentHash = value?.GetHashCode() ?? 0;

            return (seed * 23) + currentHash;
        }
    }
}
*/

// ValueObject is a base struct for value objects.
type ValueObject struct{}

// Equals checks if two ValueObject instances are equal.
func (v *ValueObject) Equals(obj interface{}) bool {
	if obj == nil || reflect.TypeOf(v) != reflect.TypeOf(obj) {
		return false
	}

	objValue := reflect.ValueOf(obj)
	vValue := reflect.ValueOf(v).Elem()

	for i := 0; i < vValue.NumField(); i++ {
		field := vValue.Field(i)
		objField := objValue.Elem().Field(i)

		if !fieldsAreEqual(field.Interface(), objField.Interface()) {
			return false
		}
	}

	return true
}

// HashCode calculates the hash code for the ValueObject.
func (v *ValueObject) HashCode() int {
	hash := 17
	vValue := reflect.ValueOf(v).Elem()

	for i := 0; i < vValue.NumField(); i++ {
		field := vValue.Field(i)
		value := field.Interface()
		hash = hashValue(hash, value)
	}

	return hash
}

func fieldsAreEqual(a, b interface{}) bool {
	if a == nil || b == nil {
		return a == b
	}

	return reflect.DeepEqual(a, b)
}

func hashValue(seed int, value interface{}) int {
	if value == nil {
		return seed * 23
	}

	currentHash := 0
	if cValue, ok := value.(ValueObject); ok {
		currentHash = cValue.HashCode()
	}

	return seed*23 + currentHash
}
