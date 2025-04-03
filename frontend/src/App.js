import React from 'react';
import './App.css';
import EventList from './components/EventList';
import CreateEvent from './components/CreateEvent';
import CreateBet from './components/CreateBet';
import BetList from './components/BetList';

function App() {
    return (
        <div className="App">
            <h1>Букмекерский сайт</h1>
            <EventList />
            <CreateEvent onEventCreated={() => window.location.reload()} />
            <CreateBet />
            <BetList />
        </div>
    );
}

export default App;