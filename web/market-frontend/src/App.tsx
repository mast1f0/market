import Navbar from "./widgets/Navbar.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from "./auth/AuthProvider.tsx";
import MainPage from "./pages/MainPage.tsx";
import AboutPage from "./pages/AboutPage.tsx";
import CategoriesPage from "./pages/CategoriesPage.tsx";
import ProfilePage from "./pages/ProfilePage.tsx";
import CartPage from "./pages/CartPage.tsx";
import AddProductPage from "./pages/AddProductPage.tsx";
import SellerPanel from "./widgets/SellerPanel.tsx";
import ProductPage from "./pages/ProductPage.tsx";
import LoginPage from "./pages/LoginPage.tsx";
import CategoryPage from "./pages/CategoryPage.tsx";

function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <div className="flex min-h-screen flex-col">
          <Navbar />
          <main className="flex-1">
            <Routes>
              <Route path="/" element={<MainPage />} />
              <Route path="/about" element={<AboutPage />} />

              <Route path="/categories" element={<CategoriesPage />} />
              <Route path="/categories/:id" element={<CategoryPage />} />

              <Route path="/profiles" element={<ProfilePage />} />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/cart" element={<CartPage />} />
              <Route path="/seller-panel" element={<SellerPanel />} />
              <Route path="/product/:id" element={<ProductPage />} />
              <Route path="/add" element={<AddProductPage />} />
            </Routes>
          </main>
          <footer className="border-t border-slate-200 bg-white py-4 text-center text-xs text-slate-500">
            Демо-витрина · {new Date().getFullYear()}
          </footer>
        </div>
      </BrowserRouter>
    </AuthProvider>
  );
}

export default App;
