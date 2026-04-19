import CategoriesGrid from "../widgets/CategoriesGrid.tsx";

const DEMO_CATEGORIES = [
  {
    id: 101,
    name: "Электроника",
    description: "Смартфоны, ноутбуки и техника для дома",
  },
  {
    id: 102,
    name: "Одежда",
    description: "Одежда и аксессуары",
  },
  {
    id: 103,
    name: "Книги",
    description: "Художественная и учебная литература",
  },
  {
    id: 104,
    name: "Дом и сад",
    description: "Мебель, декор, инструменты",
  },
];

export default function CategoriesPage() {
  return <CategoriesGrid fallbackCategories={DEMO_CATEGORIES} />;
}
