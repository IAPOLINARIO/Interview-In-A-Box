import React from 'react';
import PropTypes from "prop-types";
import {SearchOutlined} from "@material-ui/icons";

const DebouncedInput = ({ value: initialValue, onChange, debounce = 0, label = 'Search', classes, ...props}) => {
    const [value, setValue] = React.useState(initialValue.toString())

    React.useEffect(() => {
        setValue(initialValue)
    }, [initialValue])

    React.useEffect(() => {
            const timeout = setTimeout(() => {
                onChange(value)
            }, debounce)

            return () => {
                clearTimeout(timeout)
            }
        },[value])
    return (
        <>
            <label
                htmlFor="default-search"
                className="mb-2 text-sm font-medium text-gray-900 sr-only"
            >
                {label}
            </label>
            <div className="relative">
                <div className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                    <SearchOutlined />
                </div>
                <input
                    type="search"
                    className={`block p-1.5 pl-10 text-sm text-gray-900 border border-1 border-gray-400 rounded-lg bg-white
                      focus:border-fuchsia-800 focus:outline-none ${classes}`}
                    id="default-search"
                    {...props}
                    value={value}
                    onChange={(e) => {
                        setValue(e.target.value)
                    }}
                />
            </div>
        </>
    );
};

DebouncedInput.propTypes = {
    value: PropTypes.string.isRequired,
    onChange: PropTypes.func.isRequired,
    debounce: PropTypes.number,
    label: PropTypes.string,
    classes: PropTypes.string,
}

export default DebouncedInput;
