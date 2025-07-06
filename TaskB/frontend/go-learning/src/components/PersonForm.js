import React, { useState } from 'react';
import axios from 'axios';

function PersonForm() {
  const [formData, setFormData] = useState({
    name: '',
    age: '',
    phone_number: '',
    city: '',
    state: '',
    street1: '',
    street2: '',
    zip_code: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:8080/person/create', {
        ...formData,
        age: parseInt(formData.age, 10)
      });
      alert('Person created successfully!');
    } catch (err) {
      console.error('Error creating person:', err);
      alert('Error creating person');
    }
  };

  return (
    <form className="person-form" onSubmit={handleSubmit}>
      <input name="name" placeholder="Name" value={formData.name} onChange={handleChange} required />
      <input name="age" type="number" placeholder="Age" value={formData.age} onChange={handleChange} required />
      <input name="phone_number" placeholder="Phone Number" value={formData.phone_number} onChange={handleChange} required />
      <input name="city" placeholder="City" value={formData.city} onChange={handleChange} required />
      <input name="state" placeholder="State" value={formData.state} onChange={handleChange} required />
      <input name="street1" placeholder="Street 1" value={formData.street1} onChange={handleChange} required />
      <input name="street2" placeholder="Street 2" value={formData.street2} onChange={handleChange} />
      <input name="zip_code" placeholder="ZIP Code" value={formData.zip_code} onChange={handleChange} required />
      <button type="submit">Submit</button>
    </form>
  );
}

export default PersonForm;
