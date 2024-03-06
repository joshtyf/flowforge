import React from "react";

const InputField = ({ id, type, value, onChange, placeholder }) => {
  return (
    <div>
      <label htmlFor={id}>{id}</label>
      <input
        className="p-2 border border-gray-300 rounded-lg mb-4 focus:outline-none focus:border-gray-600 text-black"
        id={id}
        type={type}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
      />
    </div>
  );
};

export default InputField;
