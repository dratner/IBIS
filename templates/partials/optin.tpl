<p class="text-blue-700 mb-6 text-center title">
    Opt In
</p>
<p class="text-blue-700 text-sm mb-6 text-center">
    Thank you for volunteering for the IBIS Project to help injured birds. When you submit this form you will be entered into our volunteer database and you'll receive SMS notifications of birds you can assist.
</p>
<form class="space-y-4" method="POST" action="/optin_process">
<div>
<label for="name" class="block text-sm font-medium text-blue-700">Name</label>
<input type="tel" id="phone" name="name" required 
    placeholder="Jane Doe"
    class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50">
</div>
<div>
<label for="phone" class="block text-sm font-medium text-blue-700">Phone Number</label>
<input type="tel" id="phone" name="phone" required 
    placeholder="###-###-####"
    class="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-500 focus:ring-opacity-50">
</div>
<div class="flex items-start">
    <div class="flex items-center h-5">
        <input id="agree" name="agree" type="checkbox" required
            class="focus:ring-blue-500 h-4 w-4 text-blue-600 border border-gray-300 rounded">
    </div>
    <div class="ml-3 text-sm">
        <label for="agree" class="font-medium text-blue-700">I agree to be contacted by IBIS (operated by Snapdragon Partners.)</label>
    </div>
</div>
<div>
    <button type="submit" id="submitButton" disabled
        class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-300 cursor-not-allowed">
    Submit
    </button>
</div>
</form>
       

   
