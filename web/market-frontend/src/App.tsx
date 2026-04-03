import Navbar from "./widgets/Navbar.tsx";
import {BrowserRouter, Routes, Route} from "react-router-dom"
import MainPage from "./pages/MainPage.tsx";
import AboutPage from "./pages/AboutPage.tsx";
import CategoriesPage from "./pages/CategoriesPage.tsx";
import ProfilePage from "./pages/ProfilePage.tsx";
import CartPage from "./pages/CartPage.tsx";
import AddProductPage from "./pages/AddProductPage.tsx";
import SellerPanel from "./widgets/SellerPanel.tsx";
function App() {


  return (
      <BrowserRouter>
        <Navbar/>
        <Routes>
          <Route path="/" element={<MainPage/>} />
          <Route path="/about" element={<AboutPage/>} />
          <Route path="/categories" element={<CategoriesPage/>} />
          <Route path="/profiles" element={<ProfilePage/>} />
            <Route path="/cart" element={<CartPage/>}/>
            <Route path="/seller-panel" element={<SellerPanel/>}/>

        {/*    Для теста*/}
            <Route path={"/add"} element={<AddProductPage/>} />
        </Routes>
      </BrowserRouter>
  )
}

export default App
