import React from 'react';
import shiftLogo from "../../assets/shift.svg";

const Footer = () => {
    return (
        <div className="flex justify-between items-center absolute -bottom-1 py-4 left-60 px-6 right-0 z-50 bg-white border-t-2 border-gray-200">
            <div className="flex space-x-2">
                <p className="text-sm">Powered by
                </p>
                <img src={shiftLogo} alt="shift cars" width="45px"/>
            </div>
            <div className="text-xs text-gray-600 space-x-1">
                <span>Copyright &copy;</span><span>2023</span>
            </div>
        </div>
    );
};

export default Footer;
