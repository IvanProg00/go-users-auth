export default function Form() {
	const inputNames: string[] = ['Username', 'Password']

	const formFields = inputNames.map((v: string) => {
		return (
		<label className="block mt-3">
			<span className="text-gray-700 ml-2">{v}</span>
			<input
				type="text"
				className="mt-1 px-3 py-1 block w-full rounded-md bg-gray-100 border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 focus:outline-none"
				placeholder={v}
			/>
		</label>
	);

	})

	return (
		<div className="max-w-sm mx-auto">
			<h2 className="text-5xl text-primary">Form</h2>
			<form className="mt-5">
				{ formFields }
			</form>
		</div>
	);
}
