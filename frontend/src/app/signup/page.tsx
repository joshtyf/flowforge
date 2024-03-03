import Link from "next/link";
import React, { useState } from "react";
import axios from "axios";
import { toast } from "react-hot-toast";
import SignupForm from "../components/form/signupForm";
import PageLayout from "../components/layout/pageLayout";
import InputField from "../components/ui/inputField";

export default function SignupPage() {
  const [loading, setLoading] = useState(false);

  const onSignup = async (user) => {
    try {
      setLoading(true);
      const response = await axios.post("/api/users/signup", user);
      console.log("Signup success", response.data);
      toast.success("Signup success");
      // router.push("/login");
    } catch (error) {
      console.log("Signup failed", error.message);
      toast.error(error.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <PageLayout title={loading ? "Processing" : "Signup"} loading={loading}>
      <SignupForm onSignup={onSignup} />
      <Link href="/login">Visit login page</Link>
    </PageLayout>
  );
}
