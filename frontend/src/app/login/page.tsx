import Link from "next/link";
import React, { useState } from "react";
import axios from "axios";
import { toast } from "react-hot-toast";
import LoginForm from "../components/form/loginForm";
import PageLayout from "../components/layout/pageLayout";
import InputField from "../components/ui/inputField";

export default function LoginPage() {
  const [loading, setLoading] = useState(false);

  const onLogin = async (user) => {
    try {
      setLoading(true);
      const response = await axios.post("/api/users/login", user);
      console.log("Login success", response.data);
      toast.success("Login success");
      // router.push("/profile");
    } catch (error) {
      console.log("Login failed", error.message);
      toast.error(error.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <PageLayout title={loading ? "Processing" : "Login"} loading={loading}>
      <LoginForm onLogin={onLogin} />
      <InputField
        id="email"
        type="text"
        value={user.email}
        onChange={(value) => setUser({ ...user, email: value })}
        placeholder="email"
      />
      <InputField
        id="password"
        type="password"
        value={user.password}
        onChange={(value) => setUser({ ...user, password: value })}
        placeholder="password"
      />
      <button
        onClick={onLogin}
        className="p-2 border border-gray-300 rounded-lg mb-4 focus:outline-none focus:border-gray-600"
      >
        Login here
      </button>
      <Link href="/signup">Visit Signup page</Link>
    </PageLayout>
  );
}
