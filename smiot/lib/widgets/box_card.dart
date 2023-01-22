import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/bloc/user_bloc.dart';
import 'package:smiot/models/boxes_models.dart';
import 'package:smiot/models/user_models.dart';
import 'package:smiot/screens/dashboard_screen.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import 'package:http/http.dart' as http;

const _storage = FlutterSecureStorage();

Future<Users> fetchUser(String userId) async {
  final token = await _storage.read(key: 'accessToken');

  final response = await http.get(
      Uri.parse('https://api.rocket-translate.com/user/$userId'),
      headers: <String, String>{
        'context-type': 'application/json',
        'Authorization': 'Bearer $token'
      });

  if (response.statusCode == 200) {
    // If the server did return a 200 OK response,
    // then parse the JSON.
    return Users.fromJson(jsonDecode(response.body));
  } else {
    // If the server did not return a 200 OK response,
    // then throw an exception.
    throw Exception('Failed to load album');
  }
}

class BoxCard extends StatefulWidget {
  final Boxes box;

  const BoxCard({
    Key? key,
    required this.box,
  }) : super(key: key);

  @override
  State<BoxCard> createState() => _BoxCardState();
}

class _BoxCardState extends State<BoxCard> {
  late Future<Users> futureUser;

  @override
  void initState() {
    super.initState();
    futureUser = fetchUser(widget.box.ownerId);
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
        onTap: () {
          Navigator.of(context).push(
            MaterialPageRoute(
              builder: (context) => DashboardScreen(
                boxId: widget.box.id,
                canSub: widget.box.canSub.canSubscribe,
              ),
            ),
          );
        },
        child: Padding(
            padding: const EdgeInsets.only(top: 8, bottom: 8),
            child: FutureBuilder<Users>(
              future: futureUser,
              builder: (context, snapshot) {
                if (snapshot.data != null) {
                  return Stack(
                    alignment: const Alignment(.9, .9),
                    children: [
                      Container(
                        margin: const EdgeInsets.only(bottom: 20),
                        padding:
                            const EdgeInsets.only(top: 10, left: 20, right: 20),
                        height: 130.0,
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(15),
                          color: Colors.grey[800],
                        ),
                        width: double.infinity,
                        child: textInBox(context, snapshot.data!),
                      ),
                      CircleAvatar(
                        backgroundImage: NetworkImage(
                          'https://api.rocket-translate.com${snapshot.data!.avatar.url}',
                        ),
                        radius: 40,
                        backgroundColor: Colors.white,
                      ),
                    ],
                  );
                }
                return const CircularProgressIndicator();
              },
            )));
  }

  Column textInBox(BuildContext context, Users state) {
    // ignore: prefer_const_literals_to_create_immutables
    return Column(children: [
      Align(
        alignment: Alignment.centerLeft,
        child: Text(
          widget.box.name,
          style: const TextStyle(
            fontFamily: 'roboto',
            fontSize: 22,
            fontWeight: FontWeight.bold,
            letterSpacing: 1.8,
            color: Colors.white70,
          ),
        ),
      ),
      Padding(
        padding: const EdgeInsets.only(top: 4.0),
        child: Align(
          alignment: Alignment.centerLeft,
          child: Text(
            'สมาชิก ${widget.box.members.length}',
            style: const TextStyle(
              fontFamily: 'roboto',
              fontSize: 14,
              fontWeight: FontWeight.w200,
              letterSpacing: 1.8,
              color: Colors.white70,
            ),
          ),
        ),
      ),
      Padding(
        padding: const EdgeInsets.only(top: 38.0),
        child: Align(
          alignment: Alignment.centerLeft,
          child: Text(
            '${state.firstName} ${state.lastName}',
            style: const TextStyle(
              fontFamily: 'roboto',
              fontSize: 14,
              fontWeight: FontWeight.w200,
              letterSpacing: 1.8,
              color: Colors.white,
            ),
          ),
        ),
      )
    ]);
  }
}



/*
 BlocBuilder<UserBloc, MyState>(
            builder: (context, state) {
              if (state is StateLoading) {
                return Stack(
                  alignment: const Alignment(.9, .9),
                  children: [
                    Container(
                      margin: const EdgeInsets.only(bottom: 20),
                      padding:
                          const EdgeInsets.only(top: 10, left: 20, right: 20),
                      height: 130.0,
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(15),
                        color: Colors.grey[800],
                      ),
                      width: double.infinity,
                      child: textInBox(context, state),
                    ),
                    const CircleAvatar(
                      backgroundImage: NetworkImage(
                        'https://api.rocket-translate.com/asset/images/65c57441-7037-418e-968d-d4b4ab52e37f.png',
                      ),
                      radius: 40,
                      backgroundColor: Colors.white,
                    ),
                  ],
                );
              }
              if (state is GetUserStateSuccess) {
                
              }
              return const SizedBox();
            },
          ), */