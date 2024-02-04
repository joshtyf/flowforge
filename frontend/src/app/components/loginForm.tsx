// Importing React libraries
import React, { useState, FormEvent } from 'react';

// Import custom UI components
import Card from '../src/app/components/ui/card';
import Button from '../src/app/components/ui/button';
import { Form, FormItem, FormField, FormLabel, FormControl, FormDescription, FormMessage } from '../src/app/components/ui/form';

// Define the LoginForm functional component
const LoginForm: React.FC = () => {

  // State hooks
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  // Handle form submission
  const handleSubmit = async (e: FormEvent) => {

    // Preventing the default form submission behavior
    e.preventDefault();

    // Add subsequent logic to handle form submission here (API req...)
  };

  // Rendering the LoginForm component
  return (
    // Wrapping the form in Card component for styling
    <Card>
      {/* Form element with an onSubmit event handler */}
      <Form onSubmit={handleSubmit}>
        {/* FormItem for grouping form elements */}
        <FormItem>
          {/* FormLabel for a form field with a FormField inside */}
          <FormLabel>
            Username:
            {/* FormField to integrate with react-hook-form */}
            <FormField
              name="username"
              controlProps={{ value: username, onChange: (value: string) => setUsername(value) }}
            />
          </FormLabel>

          {/* FormLabel for another form field with a FormField inside */}
          <FormLabel>
            Password:
            {/* FormField for the password field */}
            <FormField
              name="password"
              controlProps={{ value: password, onChange: (value: string) => setPassword(value) }}
              type="password"
            />
          </FormLabel>

          {/* FormControl for the submit button */}
          <FormControl>
            <Button type="submit">Login</Button>
          </FormControl>

          {/* Example usage of FormDescription and FormMessage */}
          <FormDescription>Enter your credentials to log in.</FormDescription>
          <FormMessage>This is an error message.</FormMessage>
        </FormItem>
      </Form>
    </Card>
  );
};

export default LoginForm;
