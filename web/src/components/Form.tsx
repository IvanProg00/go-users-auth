import React from "react";
import { useState } from "react";
import { InputField } from "../util/interfaces";

export default function Form() {
	let [username, setUsername] = useState("");
	let [password, setPassword] = useState("");
	const inputNames: InputField[] = [
		{
			field: "Username",
			value: username,
			type: 'text',
			setState: setUsername,
		},
		{
			field: "Password",
			value: password,
			type: 'password',
			setState: setPassword,
		},
	];

	const formFields = inputNames.map((elem: InputField, i: number) => {
		return (
			<label className="block mt-3" key={i}>
				<span className="text-gray-700 ml-2">{elem.field}</span>
				<input
					value={elem.value}
					onChange={(e) => elem.setState(e.target.value)}
					type={elem.type}
					className="mt-1 px-3 py-1 block w-full rounded-input bg-input shadow-sm focus:ring focus:ring-input focus:outline-none"
					placeholder={elem.field}
				/>
			</label>
		);
	});

	const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();
		console.log("Form:", username, password);
	}

	return (
		<div className="max-w-sm mx-auto">
			<h2 className="text-5xl text-primary">Form</h2>
			<form className="mt-5" onSubmit={onSubmit}>
				{formFields}
				<button className="block bg-button text-button rounded-input hover:bg-blue-700 text-white font-bold py-2 px-4 focus:ring focus:ring-button focus:outline-none mt-5 mx-auto">Button</button>
			</form>
		</div>
	);
}
