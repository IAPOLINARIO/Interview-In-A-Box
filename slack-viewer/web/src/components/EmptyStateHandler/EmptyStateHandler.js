import React, {useEffect} from "react";
import {InfoOutlined} from "@material-ui/icons";

const EmptyStateBox = ({ keyword, errorMessage }) => {
    
    const message = errorMessage ?? `No ${keyword ?? ''} data currently`;
    return (
        <div className="px-3 py-12 min-h-32 text-gray-500 text-center">
            <div className="bg-fuchsia-50/50 max-w-2xl mx-auto border border-fuchsia-800/30 text-fuchsia-900/75 px-4 py-2 w-full flex flex-col space-y-3 text-sm rounded-md">
                <div className="flex items-center justify-center font-semibold space-x-3">
                    <InfoOutlined />
                    <p>{message}</p>
                </div>
                <div>
                    <p className="text-center">Use search bar and date range picker to get the results!</p>
                </div>
            </div>
        </div>
    );
};

const EmptyStateHandler = ({ children, data, errorMessage, keyword }) => {
    useEffect(() => {}, [data]);

    if ((data != null) && Array.isArray(data) && (data.length === 0)) return <EmptyStateBox keyword={keyword} errorMessage={errorMessage} />;
    if ((data != null) && typeof data === 'object' && (Object.keys(data).length === 0)) return <EmptyStateBox keyword={keyword} errorMessage={errorMessage} />;
    if (data == null) return <EmptyStateBox/>;
    return <>{children}</>;
};

export default EmptyStateHandler;
