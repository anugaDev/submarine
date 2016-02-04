package main

import (
	battleLogic "app/battle"
	"app/currentmillis"
	"app/typhenapi/type/submarine"
	"app/typhenapi/type/submarine/battle"
	webapi "app/typhenapi/web/submarine"
	"fmt"
)

// Room represents a network group for battle.
type Room struct {
	id           int64
	webAPI       *webapi.WebAPI
	info         *battle.Room
	sessions     map[int64]*Session
	battle       *battleLogic.Battle
	closeHandler func(*Room)
	join         chan *Session
	leave        chan *Session
	close        chan struct{}
}

func newRoom(id int64) (*Room, error) {
	webAPI := NewWebAPI("http://localhost:3000")

	// TODO: Validation for creatable the room in the battle server.
	res, err := webAPI.Battle.FindRoom(id)
	if err != nil {
		return nil, err
	}
	if res.Room == nil {
		return nil, fmt.Errorf("No room(%v) found.", id)
	}

	room := &Room{
		id:       id,
		webAPI:   webAPI,
		info:     res.Room,
		sessions: make(map[int64]*Session),
		battle:   battleLogic.New(currentmillis.Second * 60),
		join:     make(chan *Session, 4),
		leave:    make(chan *Session, 4),
		close:    make(chan struct{}),
	}

	go room.run()
	return room, nil
}

func (r *Room) run() {
	Log.Infof("Room(%v) has created", r.id)

loop:
	for {
		select {
		case session := <-r.join:
			r._join(session)
			r.broadcastRoom()
			session.synchronizeTime()
		case session := <-r.leave:
			r._leave(session)
			r.broadcastRoom()
		case <-r.close:
			r._close()
			break loop
		case output := <-r.battle.Gateway.Output:
			r.onBattleOutputReceive(output)
		}
	}

	if r.closeHandler != nil {
		r.closeHandler(r)
	}
}

func (r *Room) toRoomAPIType() *submarine.Room {
	members := make([]*submarine.User, len(r.sessions))
	i := 0
	for _, s := range r.sessions {
		members[i] = s.toUserAPIType()
		i++
	}
	return &submarine.Room{Id: r.id, Members: members}
}

func (r *Room) broadcastRoom() {
	typhenType := r.toRoomAPIType()
	for _, s := range r.sessions {
		s.api.Battle.SendRoom(typhenType)
	}
}

func (r *Room) _join(session *Session) {
	Log.Infof("Session(%v) has joined into Room(%v)", session.id, r.id)
	r.sessions[session.id] = session
	session.room = r
	session.disconnectHandler = func(session *Session) {
		r.leave <- session
	}

	// TODO: Add relevant room members counting.
	if len(r.sessions) >= 1 {
		r.battle.Start()
	}
}

func (r *Room) _leave(session *Session) {
	Log.Infof("Session(%v) has leaved from Room(%v)", session.id, r.id)
	session.disconnectHandler = nil
	session.room = nil
	delete(r.sessions, session.id)
}

func (r *Room) _close() {
	Log.Infof("Room(%v) has closed", r.id)
	r.battle.Gateway.Close <- struct{}{}
	for _, session := range r.sessions {
		r._leave(session)
		session.close()
	}
}

func (r *Room) onBattleOutputReceive(output interface{}) {
	switch message := output.(type) {
	case *battle.Start:
		Log.Infof("Room(%v)'s battle has started", r.id)
		for _, s := range r.sessions {
			s.api.Battle.SendStart(message)
		}
	case *battle.Finish:
		Log.Infof("Room(%v)'s battle has finished", r.id)
		for _, s := range r.sessions {
			s.api.Battle.SendFinish(message)
		}
		r.close <- struct{}{}
	}
}
