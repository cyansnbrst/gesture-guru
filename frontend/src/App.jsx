import { Routes, Route } from 'react-router-dom';
import Header from './components/Common/Header';
import Login from './components/Auth/Login';
import Register from './components/Auth/Register';
import ProtectedRoute from './components/Common/ProtectedRoute';
import GesturesList from './components/Gestures/GesturesList';
import GestureDetail from './components/Gestures/GestureDetail';

function App() {
    return (
        <div className="app">
            <Header />
            <main className="main-content">
                <Routes>
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />

                    <Route element={<ProtectedRoute />}>
                        <Route path="/gestures" element={<GesturesList />} />
                        <Route path="/gestures/:id" element={<GestureDetail />} />
                    </Route>
                </Routes>
            </main>
        </div>
    );
}

export default App;