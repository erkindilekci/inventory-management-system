import { ChangeEvent, useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { GrEdit } from "react-icons/gr";
import { RiDeleteBin6Line } from "react-icons/ri";

interface Product {
    id: number;
    name: string;
    price: number;
    quantity: number;
    category: string;
}

export const Dashboard = () => {
    const navigate = useNavigate();
    const [products, setProducts] = useState<Product[]>([]);
    const [categories, setCategories] = useState<string[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<string>("");

    const handleLogout = async () => {
        try {
            const response = await fetch("/api/logout", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                }
            });

            if (response.ok) {
                navigate("/");
            } else {
                console.error("Logout failed");
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };

    useEffect(() => {
        const fetchProducts = async () => {
            try {
                const response = await fetch("/api/products");
                const data = await response.json();
                setProducts(() => data);
            } catch (error) {
                console.error("Error fetching products:", error);
            }
        };

        fetchProducts();
    }, []);

    useEffect(() => {
        setCategories(() => Array.from(new Set(products.map(p => p.category))));
    }, [products]);

    const handleCategoryChange = (event: ChangeEvent<HTMLSelectElement>) => {
        const category = event.target.value;
        setSelectedCategory(category);
        if (category) {
            setProducts(products.filter(product => product.category === category));
        } else {
            const fetchProducts = async () => {
                try {
                    const response = await fetch("/api/products");
                    const data = await response.json();
                    setProducts(() => data);
                } catch (error) {
                    console.error("Error fetching products:", error);
                }
            };

            fetchProducts();
        }
    };

    const handleUpdate = (product: Product) => {
        navigate("/update-product", { state: { product } });
    };

    const handleDelete = async (id: number) => {
        try {
            const response = await fetch(`/api/products/${id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json"
                }
            });

            if (response.ok) {
                setProducts(prevProducts => prevProducts.filter(product => product.id !== id));
            } else {
                const result = await response.json();
                console.error("Delete failed:", result.error_message);
            }
        } catch (error) {
            console.error("Error:", error);
        }
    };


    return (
        <div className="container mx-auto px-4 sm:px-8">
            <div className="py-8">
                <div className="flex flex-row mb-1 sm:mb-0 justify-between items-center w-full">
                    <h2 className="text-2xl font-semibold leading-tight">Products</h2>
                    <button
                        onClick={handleLogout}
                        className="rounded-lg text-sm px-4 transition-all bg-black py-[0.55rem] text-white hover:bg-[#434343]">Logout
                    </button>
                </div>
                <div className="flex flex-row my-4 items-center">
                    <select
                        value={selectedCategory}
                        onChange={handleCategoryChange}
                        className="block bg-white w-full rounded-md border-0 px-2 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                    >
                        <option value="">All Categories</option>
                        {categories.map((category, index) => (
                            <option key={index} value={category}>{category}</option>
                        ))}
                    </select>
                </div>
                <div className="-mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">
                    <div className="inline-block min-w-full shadow-md rounded-lg overflow-hidden">
                        <table className="min-w-full leading-normal">
                            <thead>
                            <tr>
                                <th className="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                                    Name
                                </th>
                                <th className="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                                    Price
                                </th>
                                <th className="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                                    Quantity
                                </th>
                                <th className="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                                    Total Price
                                </th>
                                <th className="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                                    Category
                                </th>
                                <th className="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider">
                                    Actions
                                </th>
                            </tr>
                            </thead>
                            <tbody>
                            {products.map((product, index) => (
                                <tr key={index}>
                                    <td className="px-5 py-5 border-b text-left border-gray-200 bg-white text-sm">
                                        <p className="text-gray-900 whitespace-no-wrap">{product.name}</p>
                                    </td>
                                    <td className="px-5 py-5 border-b text-left border-gray-200 bg-white text-sm">
                                        <p className="text-gray-900 whitespace-no-wrap">
                                            {new Intl.NumberFormat('en-US', {
                                                style: 'currency',
                                                currency: 'USD'
                                            }).format(product.price)}
                                        </p>
                                    </td>
                                    <td className="px-5 py-5 border-b text-left border-gray-200 bg-white text-sm">
                                        <p className="text-gray-900 whitespace-no-wrap">{product.quantity}</p>
                                    </td>
                                    <td className="px-5 py-5 border-b text-left border-gray-200 bg-white text-sm">
                                        <p className="text-gray-900 whitespace-no-wrap">
                                            {new Intl.NumberFormat('en-US', {
                                                style: 'currency',
                                                currency: 'USD'
                                            }).format(parseFloat((product.quantity * product.price).toFixed(2)))}
                                        </p>
                                    </td>
                                    <td className="px-5 py-5 border-b text-left border-gray-200 bg-white text-sm">
                                        <p className="text-gray-900 whitespace-no-wrap">{product.category}</p>
                                    </td>
                                    <td className="px-5 py-5 border-b text-left border-gray-200 bg-white text-sm flex flex-row items-center">
                                        <button
                                            onClick={() => handleUpdate(product)}
                                            className="text-lg"
                                        >
                                            <GrEdit/>
                                        </button>
                                        <button
                                            onClick={() => handleDelete(product.id)}
                                            className="text-xl ml-5"
                                        >
                                            <RiDeleteBin6Line/>
                                        </button>
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    </div>
                </div>
                <div className="mt-8">
                    <Link to="/add-product"
                          className="rounded-lg border-2 border-solid bg-transparent px-4 py-2 transition-all border-black text-black hover:bg-[#dedede]"
                    >Add Product
                    </Link>
                </div>
            </div>
        </div>
    );
};
