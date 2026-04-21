import CategoriesGrid from "../widgets/CategoriesGrid.tsx";
import {useNavigate} from "react-router-dom";

export default function CategoriesPage() {
  const navigate = useNavigate();
  return (
      <CategoriesGrid onCategoryClick={(category) => {
        navigate(`/categories/${category.id}`) }} />
  );
}