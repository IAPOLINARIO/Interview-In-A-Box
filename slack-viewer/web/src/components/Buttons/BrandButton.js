import React from 'react';
import {Link} from "react-router-dom";
const BrandButton = () => {
    return (
        <Link to="/" className="flex items-center text-[20px] focus:outline-fuchsia-500">
                <div className="flex flex-col">
                    <span className="self-center font-semibold whitespace-nowrap text-gray-900">
                    Slack<span className="text-fuchsia-800">Viewer</span>
                </span>
            </div>
        </Link>
    );
};

export default BrandButton;
