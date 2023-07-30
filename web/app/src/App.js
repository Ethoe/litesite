import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from './components/pages/Home';
import Login from './components/common/Login';
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
          </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
