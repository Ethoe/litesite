import React from 'react';
import apiClient from './../../services/apiClient';
import { useNavigate, Link, NavLink } from 'react-router-dom';

const NavigationBar = ({ user, setUser }) => {
    const navigate = useNavigate();

    const onLogout = (event) => {
        event.preventDefault();

        // Send login request to the backend
        apiClient
            .get('/user/logout')
            .then((response) => response.data)
            .then((data) => {
                // Handle login response
                if (data.success) {
                    // Redirect to dashboard or home page on successful login
                    // You can use React Router to handle the navigation
                    setUser(null)
                    navigate('/'); // Replace with your desired URL
                } else {
                    console.error('Error logging out:', data.error);
                }
            })
            .catch((error) => {
                console.error('Error logging out:', error);
            });
    };


    return (
        <nav style={{ backgroundColor: 'cornflowerblue', color: 'aliceblue', padding: '10px', marginBottom: '20px' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', maxWidth: '800px', margin: '0 auto', alignItems: 'center' }}>
                <Link to="/" style={{ color: 'aliceblue', textDecoration: 'none', fontSize: '30px' }}>
                    Ethoe's Blog
                </Link>
                <ul style={{ display: 'flex', listStyle: 'none', gap: '10px', marginTop: 'revert', fontSize: '20px' }}>
                    <li>
                        <NavLink exact to="/" style={{ color: 'aliceblue', textDecoration: 'none' }} activeStyle={{ fontWeight: 'bold' }}>
                            Home
                        </NavLink>
                    </li>
                    {!user ? (
                        <li>
                            <NavLink exact to="/login" style={{ color: 'aliceblue', textDecoration: 'none' }} activeStyle={{ fontWeight: 'bold' }}>
                                Login
                            </NavLink>
                        </li>

                    ) : (
                        <>
                            <li>
                                <span style={{ color: 'aliceblue', textDecoration: 'none' }}>
                                    {user.username}
                                </span>
                            </li>
                            <li>
                                <button style={{ color: 'aliceblue', background: 'none', border: 'none', cursor: 'pointer', margin: '0', padding: '0' }} onClick={onLogout}>
                                    Logout
                                </button>
                            </li>
                        </>
                    )}
                </ul>
            </div>
        </nav>
    );
};

export default NavigationBar;
