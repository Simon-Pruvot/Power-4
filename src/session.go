package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

// Chaque visiteur a sa propre partie, retrouvée grâce à un cookie de session.
// Les parties restent en mémoire ; celles inutilisées depuis 24 h sont supprimées.

const (
	sessionCookie = "p4_session"
	sessionTTL    = 24 * time.Hour
)

type session struct {
	game     *pageData
	lastSeen time.Time
}

var (
	sessionsMu sync.Mutex
	sessions   = map[string]*session{}
)

func newGame() *pageData {
	return &pageData{joueur: []string{"/images/pion1.png", "/images/pion2.png"}, indiceJoueur: 0}
}

func newSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// visitorID renvoie l'identifiant du visiteur, en créant le cookie si besoin.
func visitorID(w http.ResponseWriter, r *http.Request) string {
	if c, err := r.Cookie(sessionCookie); err == nil && len(c.Value) == 32 {
		return c.Value
	}
	id := newSessionID()
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookie,
		Value:    id,
		Path:     "/",
		MaxAge:   int(sessionTTL.Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return id
}

// gameFor renvoie la partie du visiteur (créée à la première visite).
func gameFor(w http.ResponseWriter, r *http.Request) *pageData {
	_, game := visitor(w, r)
	return game
}

// visitor renvoie l'identifiant du visiteur et sa partie.
func visitor(w http.ResponseWriter, r *http.Request) (string, *pageData) {
	id := visitorID(w, r)
	now := time.Now()

	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	s, ok := sessions[id]
	if !ok {
		for key, old := range sessions {
			if now.Sub(old.lastSeen) > sessionTTL {
				delete(sessions, key)
			}
		}
		s = &session{game: newGame()}
		sessions[id] = s
	}
	s.lastSeen = now
	return id, s.game
}
