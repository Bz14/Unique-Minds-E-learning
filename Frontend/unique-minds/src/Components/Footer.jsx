import React from "react";
import {
  FaFacebookF,
  FaTwitter,
  FaInstagram,
  FaLinkedinIn,
  FaEnvelope,
  FaPhone,
  FaMapMarkerAlt,
} from "react-icons/fa";

const Footer = () => {
  return (
    <footer className="bg-customBlue text-white py-12 px-8 mt-48">
      <div className="container mx-auto grid grid-cols-1 md:grid-cols-3 gap-12 text-center md:text-left">
        <div>
          <h4 className="text-xl font-bold mb-4">About Us</h4>
          <p className="text-white">
            E-Learning Platform is dedicated to providing accessible and
            engaging educational content for students with Down syndrome. Our
            mission is to ensure that every learner has the opportunity to
            thrive and succeed.
          </p>
        </div>

        {/* External Links Section */}
        <div>
          <h4 className="text-xl font-bold mb-4">Quick Links</h4>
          <ul className="text-white">
            <li className="mb-3">
              <a
                href="/about"
                className="hover:text-gray-300 transition-colors"
              >
                About Us
              </a>
            </li>
            <li className="mb-3">
              <a
                href="/courses"
                className="hover:text-gray-300 transition-colors"
              >
                Courses
              </a>
            </li>
            <li className="mb-3">
              <a
                href="/donate"
                className="hover:text-gray-300 transition-colors"
              >
                Donate
              </a>
            </li>
            <li className="mb-3">
              <a
                href="/contact"
                className="hover:text-gray-300 transition-colors"
              >
                Contact Us
              </a>
            </li>
          </ul>
        </div>

        {/* Contact Section */}
        <div>
          <h4 className="text-xl font-bold mb-4">Contact Us</h4>
          <ul className="text-white">
            <li className="flex items-center justify-center md:justify-start mb-3">
              <FaEnvelope className="mr-2" />
              <a
                href="mailto:info@elearning.com"
                className="hover:text-gray-300 transition-colors"
              >
                info@elearning.com
              </a>
            </li>
            <li className="flex items-center justify-center md:justify-start mb-3">
              <FaPhone className="mr-2" />
              <a
                href="tel:+1234567890"
                className="hover:text-gray-300 transition-colors"
              >
                +1 234 567 890
              </a>
            </li>
            <li className="flex items-center justify-center md:justify-start">
              <FaMapMarkerAlt className="mr-2" />
              <span>1234 E-Learning St, Education City, Country</span>
            </li>
          </ul>
        </div>
      </div>

      <div className="container mx-auto mt-12 flex flex-col md:flex-row items-center justify-between text-center md:text-left">
        {/* Social Media Links */}
        <div className="flex justify-center mb-4 md:mb-0">
          <a
            href="https://www.facebook.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-white hover:text-gray-300 mx-3"
          >
            <FaFacebookF />
          </a>
          <a
            href="https://www.twitter.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-white hover:text-gray-300 mx-3"
          >
            <FaTwitter />
          </a>
          <a
            href="https://www.instagram.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-white hover:text-gray-300 mx-3"
          >
            <FaInstagram />
          </a>
          <a
            href="https://www.linkedin.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-white hover:text-gray-300 mx-3"
          >
            <FaLinkedinIn />
          </a>
        </div>

        {/* Copyright Section */}
        <div className="text-white text-sm">
          © 2024 E-Learning Platform. All rights reserved.
        </div>
      </div>
    </footer>
  );
};

export default Footer;
