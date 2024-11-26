"use client";

import FeaturedCourseSection from "@/app/components/Homepage/FeaturedCourseSection/featured_course_section";
import HeroSection from "@/app/components/Homepage/HeroSection/hero_section";
import BenefitsSection from "./components/Homepage/Benefits_section/benifts";

const Home = () => {
  return (
    <>
      <HeroSection />
      <FeaturedCourseSection />
      <BenefitsSection />
    </>
  );
};

export default Home;
