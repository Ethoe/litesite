import React from 'react';
import { Link, NavLink } from 'react-router-dom';

const NavigationBar = ({ user }) => {
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
                        <li>
                            <span style={{ color: 'aliceblue', textDecoration: 'none' }}>
                                {user.username}
                            </span>
                        </li>
                    )}
                </ul>
            </div>
        </nav>
    );
};

export default NavigationBar;
