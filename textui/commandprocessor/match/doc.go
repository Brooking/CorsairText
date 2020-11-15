// Package match contains the partial text matcher
//
// NewMatcher() creates a text matcher (match.Matcher), which exposes:
// Add() to add a new text entry to the matcher
// Match() which takes text and returns all of the entries for which it is a prefix
//
// There are two implementations of Matcher: the original tree based one (in tree) and a lightweight regexp based one (in regex)
package match
