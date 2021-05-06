import 'package:flutter/material.dart';
import 'package:walking_octopus/constants.dart';
import 'package:walking_octopus/size_config.dart';

import 'loginsuccess_body.dart';

class LoginSuccessScreen extends StatelessWidget {
  static String routeName = "/login_success";
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