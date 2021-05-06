import 'package:flutter/material.dart';
import '../constants.dart';
import '../size_config.dart';
import 'forgetpassword_body.dart';


class ForgetPasswordScreen extends StatelessWidget {
  static String routeName = "/forget_password";
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Walking Octopus",
          style: TextStyle(
            fontSize: getProportionateScreenWidth(24),
            color: kPrimaryColor,
            fontWeight: FontWeight.bold,
          ),),
      ),
      body: Body(),
    );
  }
}