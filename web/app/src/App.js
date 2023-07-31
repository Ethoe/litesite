import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from './components/pages/Home';
import Login from './components/common/Login';
import Register from './components/common/Register';
import NavigationBar from './components/common/NavigationBar';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <header>
          <NavigationBar />
        </header>
        <main>
          <Routes>
            <Route exact path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
          </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
