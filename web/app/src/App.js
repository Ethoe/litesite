import './App.css';
import React, { useState, useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from './components/pages/Home';
import Login from './components/common/Login';
import Register from './components/common/Register';
import NavigationBar from './components/common/NavigationBar';
import apiClient from './services/apiClient';

function App() {
  const [user, setUser] = useState(null);

  useEffect(() => {
    apiClient.get('/user/self')
      .then((response) => response.data)
      .then((data) => {
        if (data.success) {
          setUser(data.user);
        } else {
          setUser(null);
        }
      })
      .catch((error) => {
        console.error('Error checking authentication:', error);
        setUser(null);
      });
  }, []);

  return (
    <div className="App">
      <BrowserRouter>
        <header>
          <NavigationBar user={user} />
        </header>
        <main>
          <Routes>
            <Route exact path="/" element={<Home />} />
            <Route path="/login" element={<Login setUser={setUser} />} />
            <Route path="/register" element={<Register setUser={setUser} />} />
          </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
