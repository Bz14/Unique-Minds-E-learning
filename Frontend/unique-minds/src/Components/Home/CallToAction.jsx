import React from "react";
import donateImage from "../../Assets/herosecrtion.jpg";

// CallToAction component
const CallToAction = () => {
  return (
    <div className="flex flex-col md:flex-row-reverse items-center p-8 bg-customBlue rounded-lg shadow-md mt-20 mb-20">
      <div className="md:w-1/2 mb-4 md:mb-0 flex justify-center">
        <img
          src={donateImage}
          alt="Help Students"
          className="w-3/4 h-auto object-cover rounded-lg shadow-lg"
        />
      </div>
      <div className="md:w-1/2 md:pr-8 text-center md:text-left">
        <h2 className="text-3xl font-bold text-white mb-4">
          Help Students with Down Syndrome Thrive
        </h2>
        <p className="text-lg text-white mb-6">
          Your contribution can make a huge difference in the lives of students
          with Down syndrome. By donating, you're providing the resources and
          support they need to reach their full potential.
        </p>
        <a
          href="/donate"
          className="inline-block bg-blue-600 text-white text-lg font-semibold py-2 px-6 rounded-lg shadow hover:bg-blue-500 transition-colors"
        >
          Donate Now
        </a>
      </div>
    </div>
  );
};

export default CallToAction;
