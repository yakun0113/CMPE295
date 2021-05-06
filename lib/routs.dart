

import 'package:flutter/material.dart';
import 'package:walking_octopus/screens/loginsuccess_screen.dart';
import 'package:walking_octopus/screens/otp_screen.dart';
import 'package:walking_octopus/screens/signup_screen.dart';
import 'screens/forgetpassword_screen.dart';
import 'screens/launch_screen.dart';
import 'screens/signin_screen.dart';
import 'screens/signup_screen.dart';
import 'screens/complete_profile_sceen.dart';
import 'screens/otp_screen.dart';
import 'screens/home_screen.dart';


// use name routs, all routes are available herei
final Map<String, WidgetBuilder> routes = {
  LaunchScreen.routeName: (context) => LaunchScreen(),
  SignInScreen.routeName: (context) => SignInScreen(),
  ForgetPasswordScreen.routeName: (context) => ForgetPasswordScreen(),
  LoginSuccessScreen.routeName: (context) => LoginSuccessScreen(),
  SignUpScreen.routeName: (context) => SignUpScreen(),
  CompleteProfileScreen.routeName: (context) => CompleteProfileScreen(),
  OtpScreen.routeName: (context) => OtpScreen(),
  HomeScreen.routeName: (context) => HomeScreen(),
};