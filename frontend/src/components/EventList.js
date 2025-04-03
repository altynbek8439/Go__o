import React, { useState, useEffect } from 'react';
import axios from 'axios';

const EventList = () => {
    const [events, setEvents] = useState([]);

    const fetchEvents = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/v1/events');
            setEvents(response.data);
        } catch (error) {
            console.error('Ошибка при загрузке событий:', error);
        }
    };

    useEffect(() => {
        fetchEvents();
    }, []);

    return (
        <div>
            <h2>События</h2>
            <div>
                {events.map(event => (
                    <div key={event.id} className="event">
                        <strong>{event.name}</strong> (ID: {event.id})<br />
                        Дата: {new Date(event.date).toLocaleDateString()}<br />
                        Коэффициенты: Победа 1: {event.odds_win1}, Ничья: {event.odds_draw}, Победа 2: {event.odds_win2}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default EventList;