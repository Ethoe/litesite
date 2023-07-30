import React, { useState, createContext, useContext } from 'react';
import { Link, NavLink } from 'react-router-dom';
import { Dropdown } from 'react-bootstrap';

const DropdownContext = createContext();

const NavigationBar = () => {
    const [isOpen, setIsOpen] = useState(false);

    return (
        <nav style={{ backgroundColor: 'black', color: 'white', padding: '10px' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', maxWidth: '800px', margin: '0 auto' }}>
                <Link to="/" style={{ color: 'white', textDecoration: 'none' }}>
                    Ethoe's Blog
                </Link>
                <ul style={{ display: 'flex', listStyle: 'none', gap: '10px', justifyContent: 'center' }}>
                    <li>
                        <NavLink exact to="/" activeStyle={{ fontWeight: 'bold' }}>
                            Home
                        </NavLink>
                    </li>
                    <li>
                        <DropdownContext.Provider value={{ isOpen, setIsOpen }}>
                            <DropdownComponent />
                        </DropdownContext.Provider>
                    </li>
                </ul>
            </div>
        </nav>
    );
};

const DropdownComponent = () => {
    const { isOpen, setIsOpen } = useContext(DropdownContext);

    const toggleDropdown = () => {
        setIsOpen(!isOpen);
    };

    return (
        <Dropdown>
            <Dropdown.Toggle variant="link" id="account-dropdown" onClick={toggleDropdown} style={{ padding: 0, margin: 0 }}>
                Account
            </Dropdown.Toggle>
            <Dropdown.Menu show={isOpen} onClick={toggleDropdown}>
                <Dropdown.Item as={Link} to="/login">
                    Login
                </Dropdown.Item>
                {/* Add more dropdown items here as needed */}
            </Dropdown.Menu>
        </Dropdown>
    );
};

export default NavigationBar;
