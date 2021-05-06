import 'package:flutter/material.dart';
import 'package:walking_octopus/constants.dart';
import 'package:walking_octopus/size_config.dart';
import 'package:walking_octopus/screens/signup_body.dart';


class SignUpScreen extends StatelessWidget {
  static String routeName = "/sign_up";
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Walking Octopus",
          style: TextStyle(
            fontSize: getProportionateScreenWidth(24),
            color: kPrimaryColor,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
      body: Body(),
    );
  }
}