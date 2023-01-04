import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/bloc/user_bloc.dart';
import 'package:smiot/models/boxes_models.dart';
import 'package:smiot/screens/dashboard_screen.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

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
  @override
  void initState() {
    super.initState();
    context.read<UserBloc>().add(GetUserEvent(userId: widget.box.ownerId));
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
          child: BlocBuilder<UserBloc, MyState>(
            builder: (context, state) {
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
                  CircleAvatar(
                    backgroundImage: state is GetUserStateSuccess
                        ? NetworkImage(
                            'http://localhost:4000${state.user.avatar.url}',
                          )
                        : const NetworkImage(
                            'http://localhost:4000/asset/images/siot-avatar.png',
                          ),
                    radius: 40,
                    backgroundColor: Colors.white,
                  ),
                ],
              );
            },
          ),
        ));
  }

  Column textInBox(BuildContext context, MyState state) {
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
            state is GetUserStateSuccess
                ? '${state.user.firstName} ${state.user.lastName}'
                : '',
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
