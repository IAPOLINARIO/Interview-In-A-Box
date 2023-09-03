import React from 'react';
import PropTypes from 'prop-types';
import BrandButton from "../Buttons/BrandButton";
import {MenuOutlined} from "@material-ui/icons";

const SidebarToggle = ({ toggleSidebarOpen, sidebarOpen }) => {
    return (
        <div className="flex items-center justify-start">
            <button
                onClick={() => {
                    toggleSidebarOpen(!sidebarOpen)
                }}
                type="button"
                className="inline-flex items-center p-2 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200"
            >
                <span className="sr-only">Open sidebar</span>
                <MenuOutlined />
            </button>
            <BrandButton />
        </div>
    );
};

SidebarToggle.propTypes = {
    toggleSidebarOpen: PropTypes.func.isRequired,
    sidebarOpen: PropTypes.bool.isRequired,
};

export default SidebarToggle;
