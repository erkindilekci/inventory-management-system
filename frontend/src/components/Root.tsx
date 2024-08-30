import { useEffect, useState } from "react";
import Cookies from "js-cookie";
import { Link } from "react-router-dom";

export const Root = () => {
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

    useEffect(() => {
        const token = Cookies.get("token");
        if (token) {
            setIsLoggedIn(() => true);
        } else {
            setIsLoggedIn(() => false);
        }
    }, []);

    return (
        <div className="w-full h-[90vh] flex flex-col justify-between ">
            <div className="flex flex-row justify-between items-center">
                <h1 className="text-3xl font-semibold">Inventory Management System</h1>
                {isLoggedIn
                    ? <Link to="dashboard"
                            className="rounded-lg px-4 transition-all bg-black py-[0.55rem] text-white hover:bg-[#434343]">Dashboard
                    </Link>
                    : <div className="flex flex-row items-center justify-between md:space-x-6 ">
                        <Link to="signup"
                              className="rounded-lg border-2 border-solid bg-transparent px-4 py-2 transition-all border-black text-black hover:bg-[#dedede]"
                        >Sign Up
                        </Link>
                        <Link to="login"
                              className="rounded-lg px-4 transition-all bg-black py-[0.55rem] text-white hover:bg-[#434343]">Login
                        </Link>
                    </div>
                }
            </div>
            <div className="w-full flex-1 flex items-center justify-center">
                <div className="flex flex-col items-center justify-center max-h-full">
                    <h1 className="text-3xl font-semibold mb-6 italic">A modern inventory system</h1>
                    <h3 className="opacity-70">An inventory management system is a software application designed to
                        optimize the
                        tracking, organization, and control of warehouse products by providing users with a platform to
                        record and monitor product quantities, prices, and other relevant data.</h3>
                </div>
            </div>
        </div>
    );
};
