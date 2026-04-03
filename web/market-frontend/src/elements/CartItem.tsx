interface CartItemProps {
    id: number;
    name: string;
    price: number;
    quantity: number;
    image: string;
    onRemove: (id: number) => void;
    onChangeQuantity: (id: number, quantity: number) => void;
}

export default function CartItem({ id, name, price, quantity, image, onRemove, onChangeQuantity }: CartItemProps) {
    return (
        <div className="flex items-center justify-between p-4 bg-white rounded-lg shadow mb-4">
            <img src={image} alt={name} className="w-24 h-24 object-cover rounded" />

            <div className="flex-1 ml-4">
                <h3 className="text-lg font-semibold text-gray-800">{name}</h3>
                <p className="text-purple-700 font-bold mt-1">${price.toFixed(2)}</p>
            </div>

            <div className="flex items-center space-x-2">
                <button
                    className="px-2 py-1 bg-gray-200 rounded hover:bg-gray-300"
                    onClick={() => onChangeQuantity(id, quantity - 1)}
                    disabled={quantity <= 1}
                >
                    -
                </button>
                <span>{quantity}</span>
                <button
                    className="px-2 py-1 bg-gray-200 rounded hover:bg-gray-300"
                    onClick={() => onChangeQuantity(id, quantity + 1)}
                >
                    +
                </button>
            </div>

            <button
                className="ml-4 text-red-500 hover:text-red-700 font-bold"
                onClick={() => onRemove(id)}
            >
                ×
            </button>
        </div>
    );
}