type CategoryCardProps = {
    name: string;
    description: string;
    onClick?: () => void;
};

export default function CategoryCard({ name, description, onClick }: CategoryCardProps) {
    return (
        <div
            onClick={onClick}
            className="cursor-pointer bg-white rounded-lg p-4 shadow hover:shadow-lg transition-shadow duration-200 hover:scale-105"
        >
            <h2 className="text-xl font-semibold mb-2">{name}</h2>
            <p className="text-gray-600">{description}</p>
        </div>
    );
}