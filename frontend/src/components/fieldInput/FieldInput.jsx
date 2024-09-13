import React from 'react'

const FieldInput = () => {
  return (
    <div class="bg-white p-8 rounded shadow-md w-full max-w-md">
    <h2 class="text-2xl font-semibold text-gray-800 mb-6 text-center">Application Form</h2>
    <form action="#" method="post">
      <div class="mb-4">
        <label for="application" class="block text-gray-700 mb-2">Application</label>
        <input type="text" id="application" name="application" required class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"/>
      </div>
      <div class="mb-6">
        <label for="url" class="block text-gray-700 mb-2">URL</label>
        <input type="url" id="url" name="url" required class="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"/>
      </div>
      <button type="submit" class="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 transition duration-200">Submit</button>
    </form>
  </div>
  )
}

export default FieldInput
