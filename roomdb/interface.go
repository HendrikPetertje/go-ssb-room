// SPDX-License-Identifier: MIT

// Package roomdb implements all the persisted database needs of the room server.
// This includes authentication, allow/deny list managment, invite and alias creation and also the notice content for the CMS.
//
// The interfaces defined here are implemented twice. Once in SQLite for production and once as mocks for testing, generated by counterfeiter (https://github.com/maxbrunsfeld/counterfeiter).
//
// See the package documentation of roomdb/sqlite for how to update it.
// It's important not to use the types generated by sqlboiler (sqlite/models) in the argument and return values of the interfaces here.
// This would leak details of the internal implementation of the roomdb/sqlite package and we want to have full control over how these interfaces can be used.
package roomdb

import (
	"context"

	"go.mindeco.de/http/auth"
	refs "go.mindeco.de/ssb-refs"
)

type RoomConfig interface {
	GetPrivacyMode(context.Context) (PrivacyMode, error)
	SetPrivacyMode(context.Context, PrivacyMode) error
	GetDefaultLanguage(context.Context) (string, error)
	SetDefaultLanguage(context.Context, string) error
}

// AuthFallbackService allows password authentication which might be helpful for scenarios
// where one lost access to his ssb device or key.
type AuthFallbackService interface {

	// Check receives the username and password (in clear) and checks them accordingly.
	// Login might be a registered alias or a ssb id who belongs to a member.
	// If it's a valid combination it returns the user ID, or an error if they are not.
	auth.Auther

	// SetPassword creates or updates a fallback login password for this user.
	SetPassword(_ context.Context, memberID int64, password string) error

	// CreateResetToken returns a token which can be used via SetPasswordWithToken() to reset the password of a member.
	CreateResetToken(_ context.Context, createdByMember, forMember int64) (string, error)

	// SetPasswordWithToken consumes a token created with CreateResetToken() and updates the password for that member accordingly.
	SetPasswordWithToken(_ context.Context, resetToken string, password string) error
}

// AuthWithSSBService defines utility functions for the challenge/response system of sign-in with ssb
// They are particualarly of service to check valid sessions (after the client provided a solution for a challenge)
// And to log out valid sessions from the clients device.
type AuthWithSSBService interface {

	// CreateToken is used to generate a token that is stored inside a cookie.
	// It is used after a valid solution for a challenge was provided.
	CreateToken(ctx context.Context, memberID int64) (string, error)

	// CheckToken checks if the passed token is still valid and returns the member id if so
	CheckToken(ctx context.Context, token string) (int64, error)

	// RemoveToken removes a single token from the database
	RemoveToken(ctx context.Context, token string) error

	// WipeTokensForMember deletes all tokens currently held for that member
	WipeTokensForMember(ctx context.Context, memberID int64) error
}

// MembersService stores and retreives the list of internal users (members, mods and admins).
type MembersService interface {
	// Add adds a new member
	Add(_ context.Context, pubKey refs.FeedRef, r Role) (int64, error)

	// GetByID returns the member if it exists
	GetByID(context.Context, int64) (Member, error)

	// GetByFeed returns the member if it exists
	GetByFeed(context.Context, refs.FeedRef) (Member, error)

	// List returns a list of all the members.
	List(context.Context) ([]Member, error)

	// Count returns the total number of members.
	Count(context.Context) (uint, error)

	// RemoveFeed removes the feed from the list.
	RemoveFeed(context.Context, refs.FeedRef) error

	// RemoveID removes the feed for the ID from the list.
	RemoveID(context.Context, int64) error

	// SetRole changes the role of the passed member id.
	// It will return an error if the member doesn't exist.
	// It should also return an error if call would remove the last admin,
	// since only admins can change roles doing so would leave the room in a crippled state.
	SetRole(context.Context, int64, Role) error
}

// DeniedKeysService changes the lists of public keys that are not allowed to get into the room
type DeniedKeysService interface {
	// Add adds the feed to the list, together with a comment for other members
	Add(ctx context.Context, ref refs.FeedRef, comment string) error

	// HasFeed returns true if a feed is on the list.
	HasFeed(context.Context, refs.FeedRef) bool

	// HasID returns true if a member id is on the list.
	HasID(context.Context, int64) bool

	// GetByID returns the list entry for that ID or an error
	GetByID(context.Context, int64) (ListEntry, error)

	// List returns a list of all the feeds.
	List(context.Context) ([]ListEntry, error)

	// Count returns the total number of denied keys.
	Count(context.Context) (uint, error)

	// RemoveFeed removes the feed from the list.
	RemoveFeed(context.Context, refs.FeedRef) error

	// RemoveID removes the feed for the ID from the list.
	RemoveID(context.Context, int64) error
}

// AliasesService manages alias handle registration and lookup
type AliasesService interface {
	// Resolve returns all the relevant information for that alias or an error if it doesnt exist
	Resolve(context.Context, string) (Alias, error)

	// GetByID returns the alias for that ID or an error
	GetByID(context.Context, int64) (Alias, error)

	// List returns a list of all registerd aliases
	List(ctx context.Context) ([]Alias, error)

	// Register receives an alias and signature for it. Validation needs to happen before this.
	Register(ctx context.Context, alias string, userFeed refs.FeedRef, signature []byte) error

	// Revoke removes an alias from the system
	Revoke(ctx context.Context, alias string) error
}

// InvitesService manages creation and consumption of invite tokens for joining the room.
type InvitesService interface {
	// Create creates a new invite for a new member. It returns the token or an error.
	// createdBy is user ID of the admin or moderator who created it. MemberID -1 is allowed if Privacy Mode is set to Open.
	// aliasSuggestion is optional (empty string is fine) but can be used to disambiguate open invites. (See https://github.com/ssb-ngi-pointer/rooms2/issues/21)
	Create(ctx context.Context, createdBy int64) (string, error)

	// Consume checks if the passed token is still valid.
	// If it is it adds newMember to the members of the room and invalidates the token.
	// If the token isn't valid, it returns an error.
	Consume(ctx context.Context, token string, newMember refs.FeedRef) (Invite, error)

	// GetByToken returns the Invite if one for that token exists, or an error
	GetByToken(ctx context.Context, token string) (Invite, error)

	// GetByToken returns the Invite if one for that ID exists, or an error
	GetByID(ctx context.Context, id int64) (Invite, error)

	// List returns a list of all the valid invites
	List(ctx context.Context) ([]Invite, error)

	// Count returns the total number of invites.
	Count(context.Context) (uint, error)

	// Revoke removes a active invite and invalidates it for future use.
	Revoke(ctx context.Context, id int64) error
}

// PinnedNoticesService allows an admin to assign Notices to specific placeholder pages.
// like updates, privacy policy, code of conduct
type PinnedNoticesService interface {
	// List returns a list of all the pinned notices with their corresponding notices and languages
	List(context.Context) (PinnedNotices, error)

	// Set assigns a fixed page name to an page ID and a language to allow for multiple translated versions of the same page.
	Set(ctx context.Context, name PinnedNoticeName, id int64) error

	// Get returns a single notice for a name and a language
	Get(ctx context.Context, name PinnedNoticeName, language string) (*Notice, error)
}

// NoticesService is the low level store to manage single notices
type NoticesService interface {
	// GetByID returns the page for that ID or an error
	GetByID(context.Context, int64) (Notice, error)

	// Save updates the passed page or creates it if it's ID is zero
	Save(context.Context, *Notice) error

	// RemoveID removes the page for that ID.
	RemoveID(context.Context, int64) error
}

// for tests we use generated mocks from these interfaces created with https://github.com/maxbrunsfeld/counterfeiter

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/aliases.go . AliasesService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/auth.go . AuthWithSSBService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/auth_fallback.go . AuthFallbackService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/denied.go . DeniedKeysService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/fixed_pages.go . PinnedNoticesService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/invites.go . InvitesService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/members.go . MembersService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/pages.go . NoticesService

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mockdb/roomconfig.go . RoomConfig
