"use client";

import FeaturedCourseSection from "@/app/components/Homepage/FeaturedCourseSection/featured_course_section";
import HeroSection from "@/app/components/Homepage/HeroSection/hero_section";
import BenefitsSection from "./components/Homepage/Benefits_section/benifts";
import TopEducators from "./components/Homepage/Top_educators/top_educator";
import CallToAction from "./components/Homepage/Call_to_action/call_action";

const Home = () => {
  return (
    <>
      <HeroSection />
      <FeaturedCourseSection />
      <BenefitsSection />
      <TopEducators />
      <CallToAction />
    </>
  );
};

export default Home;
