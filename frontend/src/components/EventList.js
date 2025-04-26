import React, { useState, useEffect } from 'react';
import api from '../api';

const EventList = () => {
    const [events, setEvents] = useState([]);
    const [error, setError] = useState('');

    const fetchEvents = async () => {
        try {
            setError('');
            const response = await api.get('/events');
            setEvents(response.data);
        } catch (error) {
            setError('Ошибка при загрузке событий: ' + (error.response?.data?.error || error.message));
        }
    };

    useEffect(() => {
        fetchEvents();
    }, []);

    return (
        <div>
            <h2>События</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <div>
                {events.map((event) => (
                    <div key={event.id} className="event">
                        <strong>{event.name}</strong> (ID: {event.id})
                        <br />
                        Дата: {new Date(event.date).toLocaleDateString()}
                        <br />
                        Коэффициенты: Победа 1: {event.odds_win1}, Ничья: {event.odds_draw}, Победа 2: {event.odds_win2}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default EventList;