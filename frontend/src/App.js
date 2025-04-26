import React, { useContext } from 'react';
import './App.css';
import { AuthContext, AuthProvider } from './context/AuthContext';
import EventList from './components/EventList';
import CreateEvent from './components/CreateEvent';
import CreateBet from './components/CreateBet';
import BetList from './components/BetList';
import Login from './components/Login';
import Register from './components/Register';

const AppContent = () => {
    const { user, logout, loading } = useContext(AuthContext);

    if (loading) {
        return <p>Загрузка...</p>;
    }

    return (
        <div className="App">
            <h1>Букмекерский сайт</h1>
            {user ? (
                <>
                    <p>Добро пожаловать, {user.username}! Баланс: {user.balance.toFixed(2)}</p>
                    <button onClick={logout} style={{ backgroundColor: '#dc3545', marginBottom: '20px' }}>
                        Выйти
                    </button>
                    <EventList />
                    <CreateEvent onEventCreated={() => window.location.reload()} />
                    <CreateBet />
                    <BetList />
                </>
            ) : (
                <>
                    <Login />
                    <Register />
                </>
            )}
        </div>
    );
};

function App() {
    return (
        <AuthProvider>
            <AppContent />
        </AuthProvider>
    );
}

export default App;